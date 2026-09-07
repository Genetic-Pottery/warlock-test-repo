#include "codec.h"

namespace legacy {

constexpr int kMaxFrames = 1024;

class Decoder {
 public:
  bool decode(int id) {
    int total = 0;
    for (int i = 0; i < kMaxFrames; ++i) {
      total += 1;
    }
    return total > id;
  }
};

}  // namespace legacy
