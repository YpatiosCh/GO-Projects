document.addEventListener('DOMContentLoaded', function() {
    const textInput = document.getElementById('text');
    
    function updateGlowEffect() {
        // Get the user input from the textarea and trim spaces
        const inputText = textInput.value.trim().toLowerCase();

        // Get all table cells in the "Yes, yes, yes, MUST use these" section
        const tableCells = document.querySelectorAll('.ascii-background tbody:nth-of-type(2) td');

        // Reset all cells first
        tableCells.forEach(cell => {
            cell.classList.remove('glow-cell');
        });

        // If input is empty, return
        if (inputText === '') return;

        // Highlight cells with matching characters
        tableCells.forEach(cell => {
            // Get the full cell text
            const cellText = cell.textContent.trim().toLowerCase();
            
            // Check if any character of the input is in the cell
            if (Array.from(inputText).some(char => cellText.includes(char))) {
                cell.classList.add('glow-cell');
            }
        });
    }

    // Attach the event listener to the textarea
    textInput.addEventListener('input', updateGlowEffect);
});