import type { RawSymbolEntry } from "../types.ts"

export const common_symbols: Record<string, RawSymbolEntry> = {
    "©": {
        description: "Copyright sign.",
        latex: "\copyright",
        keywords: ["copyright"],
    },
    "®": {
        description: "Registered trade mark sign.",
        latex: "\textregistered",
        keywords: ["registered", "trademark"],
    },
    "™": {
        description: "Trade mark sign.",
        latex: "\texttrademark",
        keywords: ["trademark"],
    },
    "°": {
        description: "Degree sign.",
        latex: "\degree",
        keywords: ["degree"],
    },
    "•": {
        description: "Bullet point.",
        latex: "\textbullet",
        keywords: ["bullet", "point"],
    },
    "…": {
        description: "Horizontal ellipsis.",
        latex: "\dots",
        keywords: ["ellipsis", "dots"],
    },
    "–": {
        description: "En dash.",
        latex: "--",
        keywords: ["dash", "en"],
    },
    "—": {
        description: "Em dash.",
        latex: "---",
        keywords: ["dash", "em"],
    },
    "§": {
        description: "Section sign.",
        latex: "\S",
        keywords: ["section"],
    },
    "¶": {
        description: "Pilcrow sign (paragraph sign).",
        latex: "\P",
        keywords: ["paragraph", "pilcrow"],
    },
    "†": {
        description: "Dagger.",
        latex: "\dagger",
        keywords: ["dagger"],
    },
    "‡": {
        description: "Double dagger.",
        latex: "\ddagger",
        keywords: ["double", "dagger"],
    },
    "★": {
        description: "Black star.",
        latex: "\star",
        keywords: ["star", "black"],
    },
    "☆": {
        description: "White star.",
        latex: "\varstar",
        keywords: ["star", "white"],
    },
    "♠": {
        description: "Black spade suit.",
        latex: "\spadesuit",
        keywords: ["spade", "card"],
    },
    "♣": {
        description: "Black club suit.",
        latex: "\clubsuit",
        keywords: ["club", "card"],
    },
    "♥": {
        description: "Black heart suit.",
        latex: "\heartsuit",
        keywords: ["heart", "card"],
    },
    "♦": {
        description: "Black diamond suit.",
        latex: "\diamondsuit",
        keywords: ["diamond", "card"],
    },
    "✓": {
        description: "Check mark.",
        latex: "\checkmark",
        keywords: ["check", "tick", "checkmark"],
    },
    "✗": {
        description: "Ballot X.",
        latex: "\texttimes",
        keywords: ["cross", "x", "wrong"],
    },
    "✠": {
        description: "Maltese cross.",
        latex: "",
        keywords: ["cross", "maltese"],
    },
}
