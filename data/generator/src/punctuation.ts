import type { RawSymbolEntry } from "../types.ts"

export const punctuation: Record<string, RawSymbolEntry> = {
    "„": {
        description: "Double low-9 quotation mark.",
        latex: "\\quotedblbase",
        keywords: ["quote", "double", "low-9"],
    },
    "“": {
        description: "Left double quotation mark.",
        latex: "\\textquotedblleft",
        keywords: ["quote", "double", "left"],
    },
    "”": {
        description: "Right double quotation mark.",
        latex: "\\textquotedblright",
        keywords: ["quote", "double", "right"],
    },
    "‘": {
        description: "Left single quotation mark.",
        latex: "`",
        keywords: ["quote", "single", "left"],
    },
    "’": {
        description: "Right single quotation mark, apostrophe.",
        latex: "'",
        keywords: ["quote", "single", "right", "apostrophe"],
    },
    "‹": {
        description: "Single left-pointing angle quotation mark.",
        latex: "\\guillemotleft",
        keywords: ["quote", "single", "angle", "left"],
    },
    "›": {
        description: "Single right-pointing angle quotation mark.",
        latex: "\\guillemotright",
        keywords: ["quote", "single", "angle", "right"],
    },
    "«": {
        description: "Left-pointing double angle quotation mark.",
        latex: "\\guillemetleft",
        keywords: ["quote", "double", "angle", "left"],
    },
    "»": {
        description: "Right-pointing double angle quotation mark.",
        latex: "\\guillemetright",
        keywords: ["quote", "double", "angle", "right"],
    },
    "′": {
        description: "Prime.",
        latex: "\\prime",
        keywords: ["prime", "minute", "foot"],
    },
    "″": {
        description: "Double prime.",
        latex: "\\pprime",
        keywords: ["double prime", "second", "inch"],
    },
    "‴": {
        description: "Triple prime.",
        latex: "\\ppprime",
        keywords: ["triple prime", "third"],
    },
    "℗": {
        description: "Sound recording copyright sign.",
        latex: "\\textcircledP", // Requires textcomp package
        keywords: ["copyright", "sound recording"],
    },
    "№": {
        description: "Numero sign.",
        latex: "\\numero", // Requires fontawesome or similar
        keywords: ["numero", "number"],
    },
};
