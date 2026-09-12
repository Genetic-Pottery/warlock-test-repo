import { Clock } from './clock';

export type SettlementId = string & { __brand: 'settlement' };

export interface SettlementApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementApplyOptions) {}

  async apply(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementResolveOptions) {}

  async resolve(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementCompactOptions) {}

  async compact(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementValidateOptions) {}

  async validate(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementProjectOptions) {}

  async project(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementReconcileOptions) {}

  async reconcile(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementEmitOptions) {}

  async emit(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementSettleOptions) {}

  async settle(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementApplyOptions) {}

  async apply(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementResolveOptions) {}

  async resolve(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementCompactOptions) {}

  async compact(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SettlementValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SettlementValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: SettlementValidateOptions) {}

  async validate(id: SettlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SettlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
