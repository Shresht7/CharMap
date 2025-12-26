// Library
import { greek as _greek } from "./src/greek.ts";
import { math as _math } from "./src/math.ts";

const processEntry = (char: string, data: any, category: string) => {
    const unicode = `U+${char.charCodeAt(0).toString(16).padStart(4, '0')}`;
    const decimal = char.charCodeAt(0).toString();
    return { ...data, category, unicode, decimal };
};

const charmap = {
    ...Object.fromEntries(
        Object.entries(_greek).map(([char, data]) => [
            char,
            processEntry(char, data, "greek"),
        ])
    ),
    ...Object.fromEntries(
        Object.entries(_math).map(([char, data]) => [
            char,
            processEntry(char, data, "math"),
        ])
    ),
};

// Write character-map to disk as a JSON file
const outPath = "data/charmap.json";
await Deno.writeTextFile(outPath, JSON.stringify(charmap));

