<script>
	import { ENDPOINTS } from '$lib/endpoints';
	import { goto } from '$app/navigation';
	let username = '';
	let password = '';
	let errorMessage = '';

	async function handleLogin() {
		errorMessage = '';

		try {
			const response = await fetch(ENDPOINTS.loginUser, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, password }),
				credentials: 'include'
			});

			if (!response.ok) {
				throw new Error(await response.text());
			}

			goto('/');
		} catch (error) {
			console.error('Login error:', error);
			errorMessage = 'Login failed. Please try again.';
		}
	}
</script>

<body>
	<h1 class="main-header">E-Combomb</h1>
	<div class="form-container">
		<form on:submit|preventDefault={handleLogin}>
			<h1 class="form-header">Login</h1>
			<div class="form-group">
				<input id="username" type="text" placeholder="username" bind:value={username} />
			</div>

			<div class="form-group">
				<input id="password" type="password" placeholder="password" bind:value={password} />
			</div>

			{#if errorMessage}
				<div class="error-message">{errorMessage}</div>
			{/if}

			<div class="login-prompt">
				<p>Forgot your password?</p>
				<a href="/login">click here</a>
			</div>

			<button type="submit">Submit</button>
		</form>
	</div>
</body>

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

	.login-prompt {
		text-align: center;
		margin-bottom: 15px;
	}

	.login-prompt p,
	.login-prompt a {
		display: inline;
		margin-right: 5px;
	}

	.login-prompt a {
		text-decoration: none;
		color: blue;
	}

	.login-prompt a:hover {
		text-decoration: underline;
	}

	.login-prompt p {
		display: inline;
		margin-right: 5px;
	}

	.login-prompt a {
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
