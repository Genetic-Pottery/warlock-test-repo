package billing

const val LEDGER_NAME = "primary"

class Ledger(private val entries: List<Long>) {
    fun sum(): Long {
        var total = 0L
        for (entry in entries) {
            total += entry
        }
        return total
    }
}

fun newLedger(): Ledger = Ledger(emptyList())
