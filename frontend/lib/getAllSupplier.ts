import { apiKey, endPoint } from "@/constants";
import { Supplier } from "@/types";
import useSWR from "swr";

export default function getAllSupplier(token: string) {
  const fetcher = (url: string) =>
    fetch(url, {
      headers: {
        accept: "application/json",
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => {
        return res.json();
      })
      .then((json) => {
        return json.data as Supplier[];
      });
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/suppliers/all`,
    fetcher
  );

  return {
    suppliers: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
