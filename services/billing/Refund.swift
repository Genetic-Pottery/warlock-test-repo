import Foundation

public let refundWindowDays = 30

public struct Refund {
    public let amount: Int

    public func isAllowed() -> Bool {
        var total = 0
        for i in 0..<5 {
            total += i
        }
        return total > 0 && amount > 0
    }
}

public func makeRefund(amount: Int) -> Refund {
    Refund(amount: amount)
}
