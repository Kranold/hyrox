// Select the form and error message elements
const createUserForm = document.getElementById('createUserForm');
const errorMessage = document.getElementById('errorMessage');

// Add an event listener for form submission
createUserForm.addEventListener('submit', async (event) => {
  event.preventDefault(); // Prevent the default form submission behavior

  // Get the email and password values from the form
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  const username = document.getElementById('username').value;
  const fitnessGoal = document.getElementById('yourFitnessGoal').value;
  const birthday = document.getElementById('birthday').value;

  try {
    // Send a POST request to the backend login endpoint
    const response = await fetch('https://hyrox-601006340303.europe-north2.run.app/create_user', {
      method: 'POST',
      //   credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password, username, fitnessGoal, birthday }),
    });

    // Check if the login was successful
    if (response.ok) {
      const creationData = await response.json();
      console.log("userData", creationData);

    }

    // Send a POST request to the backend login endpoint
    const loginResponse = await fetch('https://hyrox-601006340303.europe-north2.run.app/login', {
      method: 'POST',
      //   credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });
    if (loginResponse.ok) {
      const loginData = await loginResponse.json();
      // Log the JWT for debugging
      // Optionally, store the JWT in localStorage or cookies
      //    localStorage.setItem('jwt', data.jwt);
      // ADD CHECK IF USER IS CONNECTED to STRAVA 
      // Redirect to another page (e.g., dashboard)
      redirectUrl = "https://hyrox-601006340303.europe-north2.run.app/strava_token_exchange" + "?token=" + loginData.jwt
      window.location.href = 'http://www.strava.com/oauth/authorize?client_id=174704&response_type=code&redirect_uri=' + redirectUrl + '&approval_prompt=force&scope=read,read_all,activity:read,profile:read_all';


    } else {
      // Handle login errors
      const errorData = await response.json();
      errorMessage.textContent = errorData.error || 'Invalid email or password';
    }
  } catch (error) {
    console.error('Error:', error);
    errorMessage.textContent = 'An error occurred. Please try again.';
  }
});