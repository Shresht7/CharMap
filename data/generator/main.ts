// Library
import { greek } from "./src/greek.ts";
import { math } from "./src/math.ts";

// Collate parts
const charmap = {
    ...greek,
    ...math,
}

// Write character-map to disk as a JSON file
const outPath = "data/charmap.json"
await Deno.writeTextFile(outPath, JSON.stringify(charmap))
