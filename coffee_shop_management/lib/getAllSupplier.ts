import { apiKey, endPoint } from "@/constants";
import { Supplier } from "@/types";
import useSWR from "swr";

const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => {
      return json.data as Supplier[];
    });

export default function getAllSupplier() {
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
