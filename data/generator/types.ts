export interface RawSymbolEntry {
    description: string;
    latex?: string;
    keywords: string[];
}

export interface SymbolEntry extends RawSymbolEntry {
    category: string;
    unicode: string;
    decimal: string;
}
