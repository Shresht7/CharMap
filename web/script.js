/** The full character map */
let CHARMAP = {}

async function loadCharacterMap() {

    let response

    try {
        response = await fetch('/charmap.json')
    } catch (e) {
        console.error(`Failed to fetch character map: ${e}`)
    }

    if (!response) {
        console.error('Failed to fetch character map')
        return
    }

    if (!response.ok) {
        console.error(`Failed to fetch character map: ${response.status}`)
        return
    }

    try {
        CHARMAP = await response.json()
    } catch (e) {
        console.error(`Failed to parse character map json: ${e}`)
    }

    console.info('Loaded character map')

    renderCharacterMap()
}

function renderCharacterMap() {
    const charMapSection = /** @type HTMLSection */ (document.getElementById('character-map'))
    if (!charMapSection) {
        console.error('Failed to find character map section')
        return
    }

    const symbolsGroupedByCategory = Object.groupBy(Object.values(CHARMAP), c => c.category)

    for (const [category, symbols] of Object.entries(symbolsGroupedByCategory)) {
        renderCharacterSection(charMapSection, category, symbols);
    }
}

function renderCharacterSection(parent, category, symbols) {
    const section = document.createElement('section');
    section.id = category;
    section.classList.add('category');

    const header = document.createElement('h2');
    header.textContent = formatDisplayString(category)
    section.appendChild(header);

    const grid = document.createElement('div');
    grid.classList.add('grid');
    section.appendChild(grid);

    for (const symbol of symbols) {
        renderSymbol(grid, symbol);
    }

    parent.appendChild(section);
}

function renderSymbol(parent, symbol) {
    const div = document.createElement('div')
    div.classList.add('symbol')
    div.innerHTML = symbol.symbol
    parent.appendChild(div)
}

document.addEventListener('DOMContentLoaded', loadCharacterMap)

// ----------------
// HELPER FUNCTIONS
// ----------------

function formatDisplayString(s) {
    const words = s.split(/[-_\s+]/g).map(capitalize)
    return words.join(' ')
}

const capitalize = (s) => s[0].toUpperCase() + s.slice(1)
