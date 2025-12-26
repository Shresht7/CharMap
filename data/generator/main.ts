// Library
import { greek as _greek } from "./src/greek.ts";
import { math as _math } from "./src/math.ts";

// Helper to add category, unicode, and decimal
const createCategoryMap = (
    rawMap: Record<string, { description: string; latex: string; keywords: string[] }>,
    category: string
) => {
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

const charmap = {
    ...createCategoryMap(_greek, "greek"),
    ...createCategoryMap(_math, "math"),
};

// Write character-map to disk as a JSON file
const outPath = "data/charmap.json";
await Deno.writeTextFile(outPath, JSON.stringify(charmap));

