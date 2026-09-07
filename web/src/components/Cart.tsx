export const CART_LIMIT = 50;

export interface CartLine {
  sku: string;
  qty: number;
}

export type CartState = "empty" | "filled";

export class Cart {
  private lines: CartLine[] = [];

  add(line: CartLine): boolean {
    for (let i = 0; i < this.lines.length; i++) {
      if (this.lines[i].sku === line.sku) return false;
    }
    this.lines.push(line);
    return true;
  }
}

export function cartState(cart: Cart): CartState {
  return "empty";
}
