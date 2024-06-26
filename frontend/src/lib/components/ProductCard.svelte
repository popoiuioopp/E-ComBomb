<script lang="ts">
	import { goto } from '$app/navigation';
	import { ENDPOINTS } from '$lib/constants/endpoints';
	import { PLACEHOLDER_IMAGE } from '$lib/constants/utils';

	export let product: {
		id: number;
		name: string;
		description: string;
		price: number;
		image: string | null;
	};

	async function addToCart() {
		try {
			const response = await fetch(ENDPOINTS.addToCart, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ product_id: product.id, quantity: 1 }),
				credentials: 'include'
			});

			if (!response.ok) {
				throw new Error('Failed to add item to cart');
			}
		} catch (error) {
			console.error('Add to cart error:', error);
		}
	}

	async function buyNow() {
		try {
			await addToCart();

			goto('/cart');
		} catch (error) {
			console.error('Buy now error:', error);
		}
	}
</script>

<div class="product-card">
	<div class="product-image">
		<img src={product.image || PLACEHOLDER_IMAGE} alt="product" />
	</div>
	<h2 class="product-name">{product.name}</h2>
	<p class="product-description">{product.description}</p>
	<p class="product-price">{product.price} baht</p>
	<div class="product-actions">
		<button on:click={addToCart} class="button-add-to-cart">Add to cart</button>
		<button on:click={buyNow} class="button-buy-now">Buy Now</button>
	</div>
</div>

<style>
	.product-card {
		width: 300px;
		border-radius: 20px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		background-color: #fff;
		padding: 20px;
		box-sizing: border-box;
		border: 1px solid #ea722f;
		margin: 20px;
	}

	.product-image {
		background-color: #eaeaea;
		height: 160px;
		border-radius: 10px;
		margin-bottom: 20px;
		overflow: hidden;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.product-image img {
		max-height: 100%;
		max-width: 100%;
		object-fit: contain;
	}

	.product-name {
		font-size: 1.5em;
		margin: 10px 0;
	}

	.product-description {
		color: #666;
		font-size: 0.9em;
		margin: 10px 0;
		overflow: hidden;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		max-width: 200px;
	}

	.product-price {
		color: #333;
		font-size: 1.2em;
		margin: 20px 0;
		font-weight: bold;
	}

	.product-actions {
		display: flex;
		justify-content: space-around;
		margin-top: 20px;
	}

	.button-add-to-cart,
	.button-buy-now {
		padding: 10px 20px;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-weight: bold;
		text-transform: uppercase;
		transition: background-color 0.3s ease;
	}

	.button-add-to-cart {
		background-color: #fff;
		color: #ea722f;
		border: 2px solid #ea722f;
	}

	.button-add-to-cart:hover {
		background-color: #f5984c;
		color: #fff;
	}

	.button-buy-now {
		background-color: #ea722f;
		color: #fff;
	}

	.button-buy-now:hover {
		background-color: #f5984c;
	}

	.product-actions button:not(:last-child) {
		margin-right: 10px;
	}
</style>
