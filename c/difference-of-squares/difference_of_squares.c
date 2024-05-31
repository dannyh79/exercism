#include "difference_of_squares.h"

static unsigned int sum_of_first_nums(unsigned int num, unsigned int acc) {
  if (num == 0) {
    return acc;
  }

  return sum_of_first_nums(num - 1, acc + num);
}

static unsigned int sum_of_first_num_squares(unsigned int num, unsigned int acc) {
  if (num == 0) {
    return acc;
  }

  return sum_of_first_num_squares(num - 1, acc + num * num);
}

unsigned int square_of_sum(unsigned int num) {
  unsigned int sum = sum_of_first_nums(num, 0);
  return sum * sum;
}

unsigned int sum_of_squares(unsigned int num) {
  return sum_of_first_num_squares(num, 0);
}

unsigned int difference_of_squares(unsigned int num) {
  return square_of_sum(num) - sum_of_squares(num);
}
