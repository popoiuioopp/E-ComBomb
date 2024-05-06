<script lang="ts">
	import { PLACEHOLDER_IMAGE } from '$lib/constants/utils';

	interface Product {
		id: number;
		name: string;
		price: number;
		description: string;
		image: string;
	}

	interface CartItem {
		Id: number;
		product_id: number;
		name: string;
		description: string;
		quantity: number;
		price: number;
		image: string | null;
		product: Product;
	}

	export let item: CartItem;
	export let removeFromCart: (productId: number) => void;
	export let updateQuantity: (productId: number, quantity: number) => void;

	function handleQuantityChange(event: Event) {
		const newQuantity = (event.target as HTMLInputElement).valueAsNumber;
		if (newQuantity >= 1) {
			updateQuantity(item.product_id, newQuantity);
		}
	}
</script>

<div class="cart-item">
	<img src={item.product.image || PLACEHOLDER_IMAGE} alt={item.name} class="cart-item-image" />
	<div class="cart-item-details">
		<h2>{item.product.name}</h2>
		<p>{item.product.description}</p>
		<p>
			Quantity: <input
				type="number"
				min="1"
				value={item.quantity}
				on:input={handleQuantityChange}
			/>
		</p>
		<p>Price: {item.product.price} baht</p>
	</div>
	<button on:click={() => removeFromCart(item.product_id)}>Remove from Cart</button>
</div>

<style>
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
