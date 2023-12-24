import { apiKey, endPoint } from "@/constants";
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
      return {
        paging: json.paging,
        data: json.data,
      };
    });

export default function getAllImportNote({
  page,
  limit,
  filterString,
}: {
  page: string;
  limit?: number;
  filterString?: string;
}) {
  const { data, error, isLoading, mutate, isValidating } = useSWR(
    `${endPoint}/importNotes?page=${page}&limit=${limit ?? "10"}${
      filterString ?? ""
    }`,
    fetcher
  );
  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
    isValidating: isValidating,
  };
}
