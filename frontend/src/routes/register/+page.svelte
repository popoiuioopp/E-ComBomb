<script>
	import { ENDPOINTS } from '$lib/endpoints';
	import { goto } from '$app/navigation';
	let username = '';
	let password = '';
	let confirmPassword = '';
	let errorMessage = '';

	const handleSubmit = async () => {
		errorMessage = '';
		if (password !== confirmPassword) {
			errorMessage = "Password doesn't match";
			return;
		}

		try {
			console.log('ENDPOINTS.registerUser', ENDPOINTS.registerUser);
			const response = await fetch(ENDPOINTS.registerUser, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username: username,
					password: password
				})
			});

			if (response.status === 302) {
				goto('/');
			} else if (response.ok) {
				const data = await response.json();
				console.log('Registration Successful', data);
				goto('/');
			} else {
				const errorData = await response.json();
				errorMessage = errorData;
				console.error('Registration failed:', errorData);
			}
		} catch (error) {
			errorMessage = 'An error occurred during registration. Please try again.';
			console.error('An error occurred during registration:', error);
		}
	};
</script>

<form on:submit|preventDefault={handleSubmit}>
	<h1 class="main-header">E-Combomb</h1>
	<div class="form-container">
		<h1 class="form-header">Register</h1>
		<div class="form-group">
			<input id="username" type="text" placeholder="username" bind:value={username} />
		</div>

		<div class="form-group">
			<input id="password" type="password" placeholder="password" bind:value={password} />
		</div>

		<div class="form-group">
			<input
				id="confirm-password"
				type="password"
				placeholder="confirm password"
				bind:value={confirmPassword}
			/>
		</div>

		{#if errorMessage}
			<div class="error-message">{errorMessage}</div>
		{/if}

		<div class="signup-prompt">
			<p>Have an account?</p>
			<a href="/login">sign in</a>
		</div>

		<button type="submit">Submit</button>
	</div>
</form>

<style>
	body {
		background-color: #ea722f;
	}

	.main-header {
		color: white;
		-webkit-text-stroke: 1px black;
		font-size: 70px;
		text-align: center;
	}

	.form-header {
		text-align: center;
		margin-top: 0px;
	}

	.form-container {
		background-color: white;
		padding: 20px;
		border-radius: 8px;
		box-shadow: 0 0 0 3px rgba(var(--primary-orange-rgb), 0.3);
		max-width: 300px;
		margin: auto;
	}

	.form-group {
		margin-bottom: 15px;
	}

	input[type='text'],
	input[type='password'] {
		width: 100%;
		padding: 8px;
		border: 1px solid #ddd;
		border-radius: 4px;
		box-sizing: border-box;
		transition:
			border-color 0.3s,
			box-shadow 0.3s;
	}

	input[type='text']:focus,
	input[type='password']:focus {
		border-color: #ea722f;
		background-color: #fffbea;
		box-shadow: 0 0 0 3px rgba(255, 127, 80, 0.3);
		outline: none;
	}

	.signup-prompt {
		text-align: center;
		margin-bottom: 15px;
	}

	.signup-prompt p,
	.signup-prompt a {
		display: inline;
		margin-right: 5px;
	}

	.signup-prompt a {
		text-decoration: none;
		color: blue;
	}

	.signup-prompt a:hover {
		text-decoration: underline;
	}

	.signup-prompt p {
		display: inline;
		margin-right: 5px;
	}

	.signup-prompt a {
		display: inline;
		text-decoration: none;
	}

	.error-message {
		text-align: center;
		color: red;
		margin-top: 10px;
		margin-bottom: 10px;
	}

	button[type='submit'] {
		width: 100%;
		padding: 10px;
		border: none;
		background-color: #ea722f;
		color: white;
		border-radius: 4px;
		cursor: pointer;
	}

	button[type='submit']:hover {
		background-color: #f5984c;
	}
</style>
