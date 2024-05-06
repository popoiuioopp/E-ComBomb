<script lang="ts">
	import { goto } from '$app/navigation';
	import User from '$lib/assets/logos/User.svelte';
	import Cart from '$lib/assets/logos/cart.svelte';

	let showDropdown = false;

	function toggleDropdown() {
		showDropdown = !showDropdown;
	}

	function handleClickOutside() {
		showDropdown = false;
	}

	function clickOutside(node: HTMLElement) {
		const handleClick = (event: MouseEvent) => {
			if (node && !node.contains(event.target as Node) && !event.defaultPrevented) {
				node.dispatchEvent(new CustomEvent('click_outside'));
			}
		};

		document.addEventListener('click', handleClick, true);
		node.addEventListener('click_outside', handleClickOutside);

		return {
			destroy() {
				document.removeEventListener('click', handleClick, true);
				node.removeEventListener('click_outside', handleClickOutside);
			}
		};
	}
</script>

<div class="navbar">
	<a href="/" class="brand">E-Combomb</a>

	<div class="logo-container">
		<div class="add-product">
			<button
				on:click={() => {
					goto('/order');
				}}
			>
				Order
			</button>
		</div>
		<div class="cart-logo">
			<a href="/cart">
				<svelte:component this={Cart} />
			</a>
		</div>

		<div class="user-logo-container" use:clickOutside>
			<button
				class="user-logo-button"
				on:click={toggleDropdown}
				aria-haspopup="true"
				aria-expanded={showDropdown}
			>
				<div class="user-logo">
					<svelte:component this={User} />
				</div>
			</button>
			{#if showDropdown}
				<div class="dropdown-menu" role="menu">
					<a href="/login" role="menuitem">Login</a>
					<a href="/register" role="menuitem">Register</a>
				</div>
			{/if}
		</div>
	</div>
</div>

<style>
	.navbar {
		display: flex;
		justify-content: space-between;
		align-items: center;

		background-color: #ea722f;
		height: 115px;
		width: 100%;
	}

	.navbar .brand {
		display: flex;
		align-items: center;

		color: white;
		text-decoration: none;
		font-weight: bold;
		-webkit-text-stroke: 1px black;
		font-size: 70px;
		text-align: center;
	}

	.logo-container {
		display: flex;
		align-items: center;
		margin-left: auto;
	}

	.navbar .user-logo,
	.navbar .cart-logo {
		display: flex;
		align-items: center;
		padding-left: 20px;
	}

	.user-logo-container {
		position: relative;
		cursor: pointer;
	}

	.dropdown-menu {
		position: absolute;
		right: 0;
		background-color: white;
		box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
		z-index: 1;
	}

	.dropdown-menu a {
		display: block;
		padding: 8px;
		text-decoration: none;
		color: black;
	}

	.dropdown-menu a:hover {
		background-color: #f1f1f1;
	}

	.user-logo-button {
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
	}

	.navbar .add-product button {
		padding: 10px 20px;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-weight: bold;
		text-transform: uppercase;
		transition: background-color 0.3s ease;
	}
</style>
