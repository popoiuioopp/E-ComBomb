<script lang="ts">
	import { goto } from '$app/navigation';
	import CartItemCard from '$lib/components/CartItemCard.svelte';
	import { ENDPOINTS } from '$lib/constants/endpoints';
	import { PLACEHOLDER_IMAGE } from '$lib/constants/utils';
	import { onMount } from 'svelte';

	interface Product {
		id: number;
		name: string;
		price: number;
		description: string;
		image: string;
	}

	interface CartItem {
		id: number;
		product_id: number;
		name: string;
		description: string;
		quantity: number;
		price: number;
		image: string | null;
		product: Product;
	}

	let cart = null;
	let cartItems: CartItem[] = [];

	$: totalPrice = cartItems.reduce((total, item) => total + item.quantity * item.product.price, 0);

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
		} catch (error) {
			console.error('Error fetching cart items:', error);
		}
	});

	async function removeFromCart(productId: number) {
		try {
			const response = await fetch(`${ENDPOINTS.deleteItemFromCart}/${productId}`, {
				method: 'DELETE',
				credentials: 'include'
			});

			if (!response.ok) {
				throw new Error('Failed to remove item from cart');
			}

			// Remove the item from the cartItems array
			cartItems = cartItems.filter((item) => item.product_id !== productId);
		} catch (error) {
			console.error('Error removing item from cart:', error);
		}
	}

	async function updateQuantity(productId: number, quantity: number) {
		try {
			console.log('requests: ', productId, quantity);
			const response = await fetch(`${ENDPOINTS.updateItemQuantity}/${productId}`, {
				method: 'PUT',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ quantity: quantity })
			});

			if (!response.ok) {
				throw new Error('Failed to update item quantity');
			}

			console.log('response: ', response);

			cartItems = cartItems.map((item) =>
				item.product_id === productId ? { ...item, quantity } : item
			);
		} catch (error) {
			console.error('Error updating item quantity:', error);
		}
	}

	async function confirmOrder() {
		try {
			const response = await fetch(ENDPOINTS.placeOder, {
				method: 'POST',
				credentials: 'include'
			});

			if (!response.ok) {
				throw new Error('Failed to place an order');
			}

			goto('/order');
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
			<CartItemCard {item} {removeFromCart} {updateQuantity} />
		{/each}
		<div class="cart-summary">
			<h3>Total Price: {totalPrice} baht</h3>
			<button on:click={confirmOrder} class="confirm-order">Confirm Order</button>
		</div>
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

	button {
		padding: 10px 20px;
		background-color: #f44336;
		color: white;
		border: none;
		border-radius: 5px;
		cursor: pointer;
	}

	.cart-summary {
		text-align: center;
		margin-top: 20px;
	}

	.confirm-order {
		padding: 10px 20px;
		background-color: #4caf50;
		color: white;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-size: 16px;
		margin-top: 10px;
	}
</style>
