// Library
import { greek } from "./src/greek.ts";

// Collate parts
const charmap = {
    ...greek
}

// Write character-map to disk as a JSON file
const outPath = "data/charmap.json"
await Deno.writeTextFile(outPath, JSON.stringify(charmap))
