import type { RawSymbolEntry } from "../types.ts"

export const scientific: Record<string, RawSymbolEntry> = {
    "°C": {
        description: "Degrees Celsius.",
        latex: "^{\\circ}\\text{C}",
        keywords: ["celsius", "temperature", "degree"],
    },
    "°F": {
        description: "Degrees Fahrenheit.",
        latex: "^{\\circ}\\text{F}",
        keywords: ["fahrenheit", "temperature", "degree"],
    },
    "ℏ": {
        description: "Planck constant over 2 pi (reduced Planck constant).",
        latex: "\\hbar",
        keywords: ["planck", "constant", "quantum", "h-bar"],
    },
    "ℓ": {
        description: "Script small l (litre).",
        latex: "\\ell",
        keywords: ["litre", "volume", "script l"],
    },
    "Å": {
        description: "Angstrom unit.",
        latex: "\\textnormal{\\AA}",
        keywords: ["angstrom", "unit", "length"],
    },
    "‰": {
        description: "Per mille sign (per thousand).",
        latex: "\\textperthousand",
        keywords: ["per mille", "per thousand", "basis points"],
    },
    " ": {
        description: "Em space.",
        latex: "\\quad",
        keywords: ["space", "em"],
    },
    "℮": {
        description: "Estimated symbol (for packaging).",
        latex: "",
        keywords: ["estimated", "estimate"],
    },
    "⇌": {
        description: "Rightwards arrow over leftwards arrow (equilibrium).",
        latex: "\\rightleftharpoons",
        keywords: ["equilibrium", "reaction", "chemical"],
    },
    "⥰": {
        description: "Reversible reaction arrow.",
        latex: "\\leftrightarrows",
        keywords: ["reversible", "reaction", "chemical"],
    },
    "⏃": {
        description: "Up-pointing triangle with dot (chemical bond).",
        latex: "",
        keywords: ["chemical bond", "bond"],
    },
};
