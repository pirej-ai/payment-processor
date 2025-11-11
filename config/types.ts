// types.ts

export type PaymentMethod = 'credit-card' | 'paypal' | 'bank-transfer';

export interface PaymentRequest {
  amount: number;
  currency: string;
  paymentMethod: PaymentMethod;
  description?: string;
}

export interface PaymentResponse {
  transactionId: string;
  amount: number;
  currency: string;
  status: 'success' | 'failure';
  description?: string;
}

export interface PaymentError {
  message: string;
  statusCode: number;
}