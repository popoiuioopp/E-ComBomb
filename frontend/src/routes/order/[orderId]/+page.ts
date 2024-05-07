import { ENDPOINTS } from '$lib/constants/endpoints';
import type { Load } from "@sveltejs/kit";

export interface Product {
  id: number;
  name: string;
  price: number;
  description: string;
  user_id: number;
  product_image: string;
}

export interface OrderItem {
  id: number;
  orderId: number;
  productId: number;
  product: Product;
  quantity: number;
}

export interface Order {
  id: number;
  status: string;
  createdAt: string;
  updatedAt: string;
  userId: number;
  items: OrderItem[];
}

export const load: Load = async ({fetch, params})=> {
  const orderId:number = parseInt(params.orderId || "")

  const fetchOrder = async (id:number) : Promise<Order>  => {
    const response = await fetch(`${ENDPOINTS.getOrderById}/${id}`)
    if (!response.ok) {
      throw new Error("Failed to fetch order")
    }
    const data = await response.json()
    return data as Order
  }

  return {
    order: await fetchOrder(orderId)
  }
}