import { apiKey, endPoint } from "@/constants";
import { Ingredient } from "@/types";
import useSWR from "swr";

export default function getAllIngredient(token: string) {
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
        return {
          data: json.data as Ingredient[],
        };
      });
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/ingredients/all`,
    fetcher
  );

  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
