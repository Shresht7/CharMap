// Library
import { greek as _greek } from "./src/greek.ts";
import { math as _math } from "./src/math.ts";

const charmap = {
    ...Object.fromEntries(
        Object.entries(_greek).map(([char, data]) => [
            char,
            { ...data, category: "greek" },
        ])
    ),
    ...Object.fromEntries(
        Object.entries(_math).map(([char, data]) => [
            char,
            { ...data, category: "math" },
        ])
    ),
};

// Write character-map to disk as a JSON file
const outPath = "data/charmap.json";
await Deno.writeTextFile(outPath, JSON.stringify(charmap));
