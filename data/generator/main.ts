// Library
import { greek as _greek } from "./src/greek.ts";
import { math as _math } from "./src/math.ts";

// Type Definitions
import type { RawSymbolEntry, SymbolEntry } from "./types.ts";

// Helper to add category, unicode, and decimal
const createCategoryMap = (
    category: string,
    rawMap: Record<string, RawSymbolEntry>,
): Record<string, SymbolEntry> => {
    return Object.fromEntries(
        Object.entries(rawMap).map(([char, data]) => {
            const unicode = `U+${char.charCodeAt(0).toString(16).padStart(4, '0')}`;
            const decimal = char.charCodeAt(0).toString();
            return [
                char,
                { ...data, category, unicode, decimal }
            ];
        })
    );
};

// The finalized character map
const charmap: Record<string, SymbolEntry> = {
    ...createCategoryMap("greek", _greek),
    ...createCategoryMap("math", _math),
};

// Write character-map to disk as a JSON file
const outPath = "data/charmap.json";
await Deno.writeTextFile(outPath, JSON.stringify(charmap));

