#include "grains.h"

uint64_t square(uint8_t index) {
  if (index < 1 || index > 64) {
    return 0;
  }

  return 1ull << (index - 1);
}

uint64_t total(void) {
  uint64_t grains = 0ull;

  for (uint8_t i = 1; i <= 64; i++) {
    grains += square(i);
  }

  return grains;
}
