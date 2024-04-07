function openNav() {
    document.getElementById("mySidenav").style.width = "360px";
}

function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
}

function filter_dropdownFunction() {
    document.getElementById("dropdown_content").classList.toggle("filter_show");
}

window.onclick = function(event) {
    if (!event.target.matches('.dropdown_button')) {
        var dropdowns = document.getElementsByClassName("dropdown_content");
        var i;
        for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('filter_show')) {
                openDropdown.classList.remove('filter_show');
            }
        }
    }
}

const team_container = document.getElementById('team_container');

// Example data to append
const newData = ["Item 4", "Item 5", "Item 6"];

// Function to append new data to the list
function appendData(data) {
    data.forEach((item, index) => {
        const newItem = document.createElement('li');
        newItem.textContent = item;
        newItem.classList.add('team_item');

        // Check if it's the first item
        if (index === 0) {
            // Create a div container for the special item
            const specialItemContainer = document.createElement('div');
            specialItemContainer.classList.add('special-item-container');
            // Append the special item to the div container
            specialItemContainer.appendChild(newItem);
            // Append the div container to the list container
            listContainer.appendChild(specialItemContainer);
        } else {
            // Append regular items directly to the list container
            listContainer.appendChild(newItem);
        }
    });
}

// Append the new data
appendData(newData);