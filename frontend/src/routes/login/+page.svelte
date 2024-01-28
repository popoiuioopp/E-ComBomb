<script>
	import { colors } from '$lib/colors';
	import { onMount } from 'svelte';
	let username = '';
	let password = '';
	let confirmPassword = '';

	const handleSubmit = async () => {};

	onMount(() => {
		document.documentElement.style.setProperty('--primary-orange-color', colors.primaryOrange);
		document.documentElement.style.setProperty('--secondary-orange-color', colors.secondaryOrange);
		document.documentElement.style.setProperty('--light-yellow-color', colors.lightYellow);
		document.documentElement.style.setProperty('--placeholder-grey-color', colors.placeholderGrey);
	});
</script>

<form on:submit={handleSubmit}>
	<h1 class="main-header">E-Combomb</h1>
	<div class="form-container">
		<form on:submit={handleSubmit}>
			<h1 class="form-header">Login</h1>
			<div class="form-group">
				<input id="username" type="text" placeholder="username" bind:value={username} />
			</div>

			<div class="form-group">
				<input id="password" type="password" placeholder="password" bind:value={password} />
			</div>

			<div class="login-prompt">
				<p>Forgot your password?</p>
				<a href="/login">click here</a>
			</div>

			<button type="submit">Submit</button>
		</form>
	</div>
</form>

<style>
	:global(body) {
		background-color: var(--primary-orange-color);
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
		border: 1px solid var(--placeholder-grey-color);
		border-radius: 4px;
		box-sizing: border-box;
		transition:
			border-color 0.3s,
			box-shadow 0.3s;
	}

	input[type='text']:focus,
	input[type='password']:focus {
		border-color: var(--primary-orange-color);
		background-color: var(--light-yellow-color);
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

	button[type='submit'] {
		width: 100%;
		padding: 10px;
		border: none;
		background-color: var(--primary-orange-color);
		color: white;
		border-radius: 4px;
		cursor: pointer;
	}

	button[type='submit']:hover {
		background-color: var(--secondary-orange-color);
	}
</style>
