import type { RawSymbolEntry } from "../types.ts"

export const arrows: Record<string, RawSymbolEntry> = {
    "←": {
        description: "Leftwards arrow.",
        latex: "\\leftarrow",
        keywords: ["arrow", "left", "west"],
    },
    "↑": {
        description: "Upwards arrow.",
        latex: "\\uparrow",
        keywords: ["arrow", "up", "north"],
    },
    "→": {
        description: "Rightwards arrow.",
        latex: "\\rightarrow",
        keywords: ["arrow", "right", "east"],
    },
    "↓": {
        description: "Downwards arrow.",
        latex: "\\downarrow",
        keywords: ["arrow", "down", "south"],
    },
    "↔": {
        description: "Left right arrow.",
        latex: "\\leftrightarrow",
        keywords: ["arrow", "left", "right", "east", "west"],
    },
    "↕": {
        description: "Up down arrow.",
        latex: "\\updownarrow",
        keywords: ["arrow", "up", "down", "north", "south"],
    },
    "⇐": {
        description: "Leftwards double arrow.",
        latex: "\\Leftarrow",
        keywords: ["arrow", "left", "double", "west"],
    },
    "⇑": {
        description: "Upwards double arrow.",
        latex: "\\Uparrow",
        keywords: ["arrow", "up", "double", "north"],
    },
    "⇒": {
        description: "Rightwards double arrow.",
        latex: "\\Rightarrow",
        keywords: ["arrow", "right", "double", "east"],
    },
    "⇓": {
        description: "Downwards double arrow.",
        latex: "\\Downarrow",
        keywords: ["arrow", "down", "double", "south"],
    },
    "⇔": {
        description: "Left right double arrow.",
        latex: "\\Leftrightarrow",
        keywords: ["arrow", "left", "right", "double", "east", "west"],
    },
    "⇕": {
        description: "Up down double arrow.",
        latex: "\\Updownarrow",
        keywords: ["arrow", "up", "down", "double", "north", "south"],
    },
}