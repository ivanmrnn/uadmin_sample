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
        team_container.appendChild(newItem); // Append directly to team_container
    });
}

// Append the new data
appendData(newData);

// Function to dynamically generate and display cards
function displayCards(n) {
    var cardGrid = document.getElementById('cardGrid');

    // Loop to create n number of card containers
    for (var i = 0; i < n; i++) {
        var cardContainer = document.createElement('div');
        cardContainer.classList.add('card_container');

        // Create a random card structure
        var cardStructure = `
            <div class="card">
                <div class="content">
                    <div class="card_front">
                        <div class="inner_card">
                            <div class="card_background_container">
                                <image src="/static/images/court.jpg" class="background_image"></image>
                            </div>
                            <div class="card_header">
                                <div class="card_name_container">
                                    <div class="card_name">
                                        <h4>Lebron James</h4>
                                    </div>
                                </div>
                                <div class="card_team_container">
                                    <div class="card_team">
                                        <image src="lakerslogo.png" class="team_logo"></image>
                                    </div>
                                </div>
                            </div>
                            <div class="card_image_container">
                                
                                <div class="card_player_container">
                                    <image src="/lebron.png" class="player_image"></image>
                                </div>
                                <div class="card_stats_container">
                                    <div class="card_stats">
                                        <div class="card_stats_title"><h5>PPG</h5></div>
                                        <div class="card_stats_value"><h4>25.5</h4></div>
                                    </div>
                                    <div class="card_stats">
                                        <div class="card_stats_title"><h5>PPG</h5></div>
                                        <div class="card_stats_value"><h4>25.5</h4></div>
                                    </div>
                                    <div class="card_stats">
                                        <div class="card_stats_title"><h5>PPG</h5></div>
                                        <div class="card_stats_value"><h4>25.5</h4></div>
                                    </div>
                                    <div class="card_stats">
                                        <div class="card_stats_title"><h5>PPG</h5></div>
                                        <div class="card_stats_value"><h4>25.5</h4></div>
                                    </div>
                                </div>
                            </div>
                            <div class="card_body">
                                <div class="card_body_row">
                                    <div class="card_info">
                                        <div class="card_info_title"><h5>Birthdate</h5></div>
                                        <div class="card_info_value"><h4>12/30/84</h4></div>
                                    </div>
                                    <div class="card_info">
                                        <div class="card_info_title"><h5>Age</h5></div>
                                        <div class="card_info_value"><h4>39 yrs</h4></div>
                                    </div>
                                    <div class="card_info">
                                        <div class="card_info_title"><h5>Experience</h5></div>
                                        <div class="card_info_value"><h4>20 yrs</h4></div>
                                    </div>
                                </div>
                                <div class="card_body_row">
                                    <div class="card_info">
                                        <div class="card_info_title"><h5>Height</h5></div>
                                        <div class="card_info_value"><h4>6'9"</h4></div>
                                    </div>
                                    <div class="card_info">
                                        <div class="card_info_title"><h5>Weight</h5></div>
                                        <div class="card_info_value"><h4>250lb</h4></div>
                                    </div>
                                    <div class="card_info">
                                        <div class="card_info_title"><h5>Country</h5></div>
                                        <div class="card_info_value"><h4>USA</h4></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="card_back">
                        <image src="/static/images/card_back.png" class="card_back_img"></image>
                    </div>
                </div>
            </div>
        `;

        // Set the innerHTML of cardContainer to the random card structure
        cardContainer.innerHTML = cardStructure;

        cardGrid.appendChild(cardContainer);
    }
}

// Call the function to display 8 cards initially
displayCards(8);