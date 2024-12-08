document.addEventListener('DOMContentLoaded', function() {
    const searchInput = document.getElementById('search');
    const suggestionBox = document.querySelector('.suggestion-box');
    let debounceTimer;

    searchInput.addEventListener('input', function() {
        clearTimeout(debounceTimer);
        const query = this.value.trim();

        if (query.length === 0) {
            suggestionBox.style.display = 'none';
            return;
        }

        // Debounce the search to avoid too many requests
        debounceTimer = setTimeout(() => {
            fetch(`/api/search?q=${encodeURIComponent(query)}`)
                .then(response => response.json())
                .then(results => {
                    suggestionBox.innerHTML = '';
                    
                    if (results.length === 0) {
                        suggestionBox.style.display = 'none';
                        return;
                    }

                    results.forEach(result => {
                        const div = document.createElement('div');
                        div.className = 'suggestion-item';
                        div.innerHTML = `
                            <span class="suggestion-text">${highlightMatch(result.text, query)}</span>
                            <span class="suggestion-type">
                                ${result.type === 'member' ? `${result.type} of ${result.context}` : 
                                  result.type === 'venue' ? `venue for ${result.context}` :
                                  result.type}
                            </span>
                        `;
                        
                            div.addEventListener('click', () => {
                                if (result.type === 'location') {
                                    window.location.href = `/location?q=${encodeURIComponent(result.originalText)}`;
                                } else if (result.type === 'concert date') {
                                    window.location.href = `/date?q=${encodeURIComponent(result.text)}`;
                                } else if (result.type === 'venue') {
                                    window.location.href = `/artist/${result.artistId}`;
                                } else if (result.type === 'member') {
                                    window.location.href = `/artist/${result.artistId}`;
                                } else {
                                    window.location.href = `/artist/${result.artistId}`;
                                }
                            });
                        
                        suggestionBox.appendChild(div);
                    });
                    
                    suggestionBox.style.display = 'block';
                })
                .catch(error => {
                    console.error('Search error:', error);
                });
        }, 300); // 300ms debounce delay
    });

    // Handle keyboard navigation
    searchInput.addEventListener('keydown', function(e) {
        const items = suggestionBox.getElementsByClassName('suggestion-item');
        const activeItem = suggestionBox.querySelector('.suggestion-item.active');
        let index = Array.from(items).indexOf(activeItem);

        if (e.key === 'ArrowDown' || e.key === 'ArrowUp') {
            e.preventDefault();

            if (e.key === 'ArrowDown') {
                index = index < items.length - 1 ? index + 1 : 0;
            } else {
                index = index > 0 ? index - 1 : items.length - 1;
            }

            if (activeItem) activeItem.classList.remove('active');
            items[index].classList.add('active');
            items[index].scrollIntoView({ block: 'nearest' });
        }

        if (e.key === 'Enter' && activeItem) {
            e.preventDefault();
            activeItem.click();
        }

        if (e.key === 'Escape') {
            suggestionBox.style.display = 'none';
            searchInput.blur();
        }
    });

    // Close suggestions when clicking outside
    document.addEventListener('click', function(e) {
        if (!suggestionBox.contains(e.target) && e.target !== searchInput) {
            suggestionBox.style.display = 'none';
        }
    });

    // Highlight matching text
    function highlightMatch(text, query) {
        const regex = new RegExp(`(${query})`, 'gi');
        return text.replace(regex, '<span class="highlight-text">$1</span>');
    }

    // Focus input when pressing '/' key
    document.addEventListener('keydown', function(e) {
        if (e.key === '/' && document.activeElement !== searchInput) {
            e.preventDefault();
            searchInput.focus();
        }
    });
});