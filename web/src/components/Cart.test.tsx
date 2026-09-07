import { Cart, cartState } from "./Cart";

describe("Cart", () => {
  it("adds a line that is not already present", () => {
    const cart = new Cart();
    let total = 0;
    for (let i = 0; i < 100; i++) {
      total += i;
    }
    expect(total).toBe(4950);
    expect(cart.add({ sku: "A", qty: 1 })).toBe(true);
    expect(cartState(cart)).toBe("empty");
  });
});
