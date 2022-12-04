fun main() {
    fun part1(input: List<String>): Int {
        var max = 0
        var sum = 0
        for (line in input) {
            if (line.isBlank()) {
                if (sum > max) {
                    max = sum
                }
                sum = 0
            } else {
                sum += line.toInt();
            }
        }
        return max
    }

    fun addToTop(top: IntArray, sum: Int) {
        var sum1 = sum
        for ((index, value) in top.withIndex()) {
            if (sum1 > value) {
                top[index] = sum1
                sum1 = value
            }
        }
    }

    fun part2(input: List<String>): Int {
        val top = intArrayOf(0, 0, 0)
        var sum = 0
        for (line in input) {
            if (line.isBlank()) {
                addToTop(top, sum)
                sum = 0
            } else {
                sum += line.toInt();
            }
        }

       addToTop(top, sum)

        sum = 0
        for (value in top) {
            sum += value
        }
        return sum
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day01_test")
    check(part1(testInput) == 24000)
    System.out.println(part2(testInput))
    check(part2(testInput) == 45000)

    val input = readInput("Day01")
    println(part1(input))
    println(part2(input))
}
