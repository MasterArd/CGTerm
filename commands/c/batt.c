#include <fcntl.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <termios.h>
#include <unistd.h>

#define BUF_SIZE 256
#define BOX_WIDTH 52

static volatile sig_atomic_t running = 1;
static struct termios orig_termios;

/* ---------------- terminal ---------------- */

void reset_terminal(void) {
    tcsetattr(STDIN_FILENO, TCSANOW, &orig_termios);
}

void handle_sigint(int sig) {
    (void)sig;
    running = 0;
}

int enable_raw_mode(void) {
    if (!isatty(STDIN_FILENO))
        return 0;

    if (tcgetattr(STDIN_FILENO, &orig_termios) == -1)
        return 0;

    atexit(reset_terminal);

    struct termios raw = orig_termios;
    raw.c_lflag &= ~(ECHO | ICANON);
    raw.c_cc[VMIN] = 0;
    raw.c_cc[VTIME] = 0;

    if (tcsetattr(STDIN_FILENO, TCSANOW, &raw) == -1)
        return 0;

    return 1;
}

/* ---------------- input ---------------- */

int read_input_char(char *out) {
    char c;
    ssize_t n = read(STDIN_FILENO, &c, 1);
    if (n == 1) {
        *out = c;
        return 1;
    }
    return 0;
}

/* ---------------- sysfs helpers ---------------- */

int read_value(const char *path, char *out, size_t size) {
    FILE *f = fopen(path, "r");
    if (!f)
        return 0;

    if (!fgets(out, size, f)) {
        fclose(f);
        return 0;
    }

    fclose(f);

    size_t len = strlen(out);
    if (len && out[len - 1] == '\n')
        out[len - 1] = '\0';

    return 1;
}

long read_long(const char *path) {
    char buf[BUF_SIZE];
    if (!read_value(path, buf, sizeof(buf)))
        return -1;
    return atol(buf);
}

/* ---------------- battery detection ---------------- */

int find_battery(char *out, size_t size) {
    DIR *d = opendir("/sys/class/power_supply");
    if (!d)
        return 0;

    struct dirent *e;

    while ((e = readdir(d)) != NULL) {
        if (e->d_name[0] == '.')
            continue;

        if (strncmp(e->d_name, "BAT", 3) == 0 ||
            strncmp(e->d_name, "CMB", 3) == 0) {
            snprintf(out, size, "%s", e->d_name);
            closedir(d);
            return 1;
        }
    }

    closedir(d);
    return 0;
}

/* ---------------- UI helpers ---------------- */

void hline(const char *l, const char *m, const char *r) {
    printf("%s", l);
    for (int i = 0; i < BOX_WIDTH - 2; i++)
        printf("%s", m);
    printf("%s\n", r);
}

void row(const char *k, const char *v) {
    printf("│ %-10s: %-36s│\n", k, v);
}

/* ---------------- battery ---------------- */

void print_battery(const char *bat) {
    char path[256];

    char vendor[BUF_SIZE] = "N/A";
    char model[BUF_SIZE] = "N/A";
    char status[BUF_SIZE] = "N/A";
    char cap[BUF_SIZE] = "N/A";

    snprintf(path, sizeof(path), "/sys/class/power_supply/%s/manufacturer", bat);
    read_value(path, vendor, sizeof(vendor));

    snprintf(path, sizeof(path), "/sys/class/power_supply/%s/model_name", bat);
    read_value(path, model, sizeof(model));

    snprintf(path, sizeof(path), "/sys/class/power_supply/%s/status", bat);
    read_value(path, status, sizeof(status));

    snprintf(path, sizeof(path), "/sys/class/power_supply/%s/capacity", bat);
    read_value(path, cap, sizeof(cap));

    long full = read_long("/sys/class/power_supply/BAT0/energy_full");
    if (full < 0)
        full = read_long("/sys/class/power_supply/BAT0/charge_full");

    long design = read_long("/sys/class/power_supply/BAT0/energy_full_design");
    if (design < 0)
        design = read_long("/sys/class/power_supply/BAT0/charge_full_design");

    double health = -1, wear = -1;

    if (full > 0 && design > 0) {
        health = ((double)full / design) * 100.0;
        wear = 100.0 - health;
    }

    hline("┌", "─", "┐");

    char title[64];
    snprintf(title, sizeof(title), "Battery Monitor (%s)", bat);
    printf("│ %-50s│\n", title);

    hline("├", "─", "┤");

    row("Battery", bat);
    row("Vendor", vendor);
    row("Model", model);
    row("Status", status);
    row("Charge", cap);

    char tmp[64];

    if (health >= 0) {
        snprintf(tmp, sizeof(tmp), "%.1f%%", health);
        row("Health", tmp);

        snprintf(tmp, sizeof(tmp), "%.1f%%", wear);
        row("Wear", tmp);
    } else {
        row("Health", "N/A");
        row("Wear", "N/A");
    }

    hline("└", "─", "┘");
    putchar('\n');
}

/* ---------------- main loop ---------------- */

void runBatt(void) {
    signal(SIGINT, handle_sigint);
    enable_raw_mode();

    char bat[64] = {0};
    if (!find_battery(bat, sizeof(bat)))
        snprintf(bat, sizeof(bat), "BAT0");

    while (running) {
        printf("\033[H\033[J");

        print_battery(bat);

        printf("Press 'q' to quit or Ctrl+C\n");
        fflush(stdout);

        for (int i = 0; i < 20 && running; i++) {
            char c;
            if (read_input_char(&c)) {
                if (c == 'q' || c == 'Q') {
                    running = 0;
                    reset_terminal();
                    break;
                }
            }
            usleep(1);
        }
    }
}
