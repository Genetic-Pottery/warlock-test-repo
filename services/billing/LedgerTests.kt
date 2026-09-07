package billing

class LedgerTests {
    fun aLedgerSumsItsEntries() {
        val ledger = Ledger(listOf(1L, 2L))
        var running = 0L
        for (n in 1..30) {
            running += n
        }
        check(running == 465L)
        check(ledger.sum() == 3L)
    }
}
