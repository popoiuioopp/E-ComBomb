<script lang="ts">
	import { goto } from '$app/navigation';
	import User from '$lib/assets/logos/User.svelte';
	import Cart from '$lib/assets/logos/cart.svelte';
	import { onDestroy } from 'svelte';
	import { authStore, logout } from '$lib/store';
	import { Roles } from '$lib/constants/roles';

	let showDropdown = false;
	let showRoleDropdown = false;

	// Subscribe to the authStore
	let isLoggedIn = false;
	let userRole = 'Guest';

	const unsubscribe = authStore.subscribe(($authStore) => {
		isLoggedIn = $authStore.isLoggedIn;
		userRole = $authStore.userRole;
	});

	onDestroy(() => {
		unsubscribe();
	});

	function toggleDropdown() {
		showDropdown = !showDropdown;
	}

	function toggleRoleDropdown() {
		showRoleDropdown = !showRoleDropdown;
	}

	function handleClickOutside() {
		showDropdown = false;
		showRoleDropdown = false;
	}

	function changeRole(role: Roles) {
		authStore.set({ userRole: role, isLoggedIn: true }); // Update role in the store
		showRoleDropdown = false;
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

	function handleLogout() {
		logout();
		goto('/');
	}
</script>

<div class="navbar">
	<a href="/" class="brand">E-Combomb</a>

	<div class="logo-container">
		{#if isLoggedIn}
			<div class="navbar-button">
				<button on:click={() => goto('/order')}> Order </button>
			</div>

			<div class="cart-logo">
				<a href="/cart">
					<svelte:component this={Cart} />
				</a>
			</div>
		{/if}

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
					{#if !isLoggedIn}
						<a href="/login" role="menuitem">Login</a>
						<a href="/register" role="menuitem">Register</a>
					{:else}
						<button on:click={handleLogout}>Logout</button>
						<button on:click={toggleRoleDropdown}>Change Role</button>
					{/if}
					{#if showRoleDropdown}
						<div class="role-dropdown-menu">
							<button on:click={() => changeRole(Roles.Buyer)}>Buyer</button>
							<button on:click={() => changeRole(Roles.Seller)}>Seller</button>
							<button on:click={() => changeRole(Roles.Delivery)}>Delivery</button>
						</div>
					{/if}
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

	.navbar .navbar-button button {
		padding: 10px 20px;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-weight: bold;
		text-transform: uppercase;
		transition: background-color 0.3s ease;
	}

	.role-dropdown-menu {
		position: absolute;
		top: 100%;
		left: 0;
		background-color: white;
		box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
		z-index: 2;
	}

	.role-dropdown-menu button {
		display: block;
		width: 100%;
		padding: 8px;
		text-align: left;
		border: none;
		background: none;
		cursor: pointer;
	}

	.role-dropdown-menu button:hover {
		background-color: #f1f1f1;
	}
</style>
