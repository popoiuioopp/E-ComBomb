<script lang="ts">
	import { ENDPOINTS } from '$lib/constants/endpoints';
	import { onMount } from 'svelte';
	let placeholderImage =
		'https://ichef.bbci.co.uk/news/976/cpsprodpb/16620/production/_91408619_55df76d5-2245-41c1-8031-07a4da3f313f.jpg';

	interface CartItem {
		id: number;
		productId: number;
		name: string;
		description: string;
		quantity: number;
		price: number;
		image: string | null;
	}

	let cart = null;
	let cartItems: CartItem[] = [];

	onMount(async () => {
		// Fetch cart items from the backend
		try {
			const response = await fetch(ENDPOINTS.getCart, {
				credentials: 'include'
			});

			if (!response.ok) {
				throw new Error('Failed to fetch cart items');
			}

			cart = await response.json();
			cartItems = cart.Items;
			console.log(cartItems);
		} catch (error) {
			console.error('Error fetching cart items:', error);
		}
	});

	async function removeFromCart(productId: number) {
		// Call backend to remove item from cart
		try {
			const response = await fetch(`http://localhost:8080/cart/${productId}`, {
				method: 'DELETE',
				credentials: 'include'
			});

			if (!response.ok) {
				throw new Error('Failed to remove item from cart');
			}

			// Remove the item from the cartItems array
			cartItems = cartItems.filter((item) => item.productId !== productId);
		} catch (error) {
			console.error('Error removing item from cart:', error);
		}
	}
</script>

<h1 class="cart-header">Shopping Cart</h1>

<hr class="divider" />

{#if cartItems.length === 0}
	<p>Your cart is empty.</p>
{:else}
	<div class="cart-items">
		{#each cartItems as item (item.id)}
			<div class="cart-item">
				<img src={item.image || placeholderImage} alt={item.name} class="cart-item-image" />
				<div class="cart-item-details">
					<h2>{item.name}</h2>
					<p>{item.description}</p>
					<p>Quantity: {item.quantity}</p>
					<p>Price: {item.price} baht</p>
				</div>
				<button on:click={() => removeFromCart(item.productId)}>Remove from Cart</button>
			</div>
		{/each}
	</div>
{/if}

<style>
	.cart-header {
		display: flex;
		justify-content: center;
	}

	.divider {
		width: 95%;
		margin: auto;
	}

	.cart-items {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.cart-item {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 10px;
		border-bottom: 1px solid #ccc;
		width: 80%;
	}

	.cart-item-image {
		width: 100px;
		height: auto;
	}

	.cart-item-details {
		flex-grow: 1;
		padding: 0 20px;
	}

	button {
		padding: 10px 20px;
		background-color: #f44336;
		color: white;
		border: none;
		border-radius: 5px;
		cursor: pointer;
	}
</style>
