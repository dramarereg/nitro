// Copyright 2024, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

use crate::evm::api::Ink;
use derivative::Derivative;
use std::time::{Duration, Instant};

// Benchmark is used to track the performance of a block of code in stylus
#[derive(Derivative, Clone, Copy)]
#[derivative(Debug)]
pub struct Benchmark {
    pub timer: Instant,
    pub elapsed: Option<Duration>,
    pub ink_start: Ink,
    pub ink_total: Option<Ink>,
}
