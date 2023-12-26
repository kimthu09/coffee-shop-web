import { apiKey, endPoint } from "@/constants";
import axios from "axios";

export default async function createImportNote({
  details,
  id,
  supplierId,
  token,
}: {
  details: {
    ingredientId: string;
    amountImport: number;
    price: number;
    isReplacePrice?: boolean;
  }[];
  id?: string;
  supplierId: string;
  token: string;
}) {
  const url = `${endPoint}/importNotes`;

  const data = {
    details,
    id: id,
    supplierId: supplierId,
  };
  console.log(data);
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,

    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
