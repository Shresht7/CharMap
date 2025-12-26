// Library
import { greek as _greek } from "./src/greek.ts";
import { math as _math } from "./src/math.ts";
import { arrows as _arrows } from "./src/arrows.ts";
import { currency as _currency } from "./src/currency.ts";
import { common_symbols as _common_symbols } from "./src/common_symbols.ts";
import { emojis as _emojis } from "./src/emojis.ts";
import { punctuation as _punctuation } from "./src/punctuation.ts";

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
    ...createCategoryMap("arrows", _arrows),
    ...createCategoryMap("currency", _currency),
    ...createCategoryMap("common_symbols", _common_symbols),
    ...createCategoryMap("emojis", _emojis),
    ...createCategoryMap("punctuation", _punctuation),
};

// Write character-map to disk as a JSON file
const outPath = "data/charmap.json";
await Deno.writeTextFile(outPath, JSON.stringify(charmap));

// Create individual category JSON files
const categories = {
    greek: _greek,
    math: _math,
    arrows: _arrows,
    currency: _currency,
    common_symbols: _common_symbols,
    emojis: _emojis,
    punctuation: _punctuation,
};

const outDir = "data/generator/out";
await Deno.mkdir(outDir, { recursive: true });

for (const categoryName in categories) {
    const categoryRawMap = categories[categoryName as keyof typeof categories];
    const categoryMap = createCategoryMap(categoryName, categoryRawMap);
    const categoryOutPath = `${outDir}/${categoryName}.json`;
    await Deno.writeTextFile(categoryOutPath, JSON.stringify(categoryMap, null, 2));
}
