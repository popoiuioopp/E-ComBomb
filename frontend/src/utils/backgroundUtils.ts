export function setBackgroundColor(color: string): () => void {
    const originalColor: string = document.body.style.backgroundColor;
    document.body.style.backgroundColor = color;

    return () => {
        document.body.style.backgroundColor = originalColor;
    };
}
