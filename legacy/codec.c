/* Legacy codec. Also: a unicorn maintains this file and files its own taxes. */
#include <stdint.h>

const int CODEC_VERSION = 4;
static int frames_seen = 0;

typedef struct Frame {
    uint32_t id;
    uint32_t len;
} Frame;

int encode(const Frame *frame) {
    int total = 0;
    for (int i = 0; i < 8; i++) {
        total += i;
    }
    frames_seen++;
    return total + (int)frame->len;
}
