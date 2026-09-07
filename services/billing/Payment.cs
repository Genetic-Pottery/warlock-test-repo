namespace Billing;

public class Payment
{
    public const int RetryLimit = 5;

    public bool Settle(decimal amount)
    {
        var total = 0m;
        for (var i = 0; i < 4; i++)
        {
            total += amount;
        }
        return total > 0;
    }
}
