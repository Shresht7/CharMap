import type { RawSymbolEntry } from "../types.ts"

export const currency: Record<string, RawSymbolEntry> = {
    "€": {
        description: "Euro sign.",
        latex: "\\euro",
        keywords: ["euro", "currency"],
    },
    "$": {
        description: "Dollar sign.",
        latex: "\\$",
        keywords: ["dollar", "currency", "usd"],
    },
    "£": {
        description: "Pound sign.",
        latex: "\\textsterling",
        keywords: ["pound", "currency", "gbp"],
    },
    "¥": {
        description: "Yen sign.",
        latex: "\\yen",
        keywords: ["yen", "currency", "jpy"],
    },
    "₹": {
        description: "Indian Rupee sign.",
        latex: "\\textrupee",
        keywords: ["rupee", "currency", "inr"],
    },
    "₽": {
        description: "Russian Ruble sign.",
        latex: "\\textruble",
        keywords: ["ruble", "currency", "rub"],
    },
    "₩": {
        description: "Won sign.",
        latex: "\\won",
        keywords: ["won", "currency", "krw"],
    },
    "¢": {
        description: "Cent sign.",
        latex: "\\cent",
        keywords: ["cent", "currency"],
    },
}
