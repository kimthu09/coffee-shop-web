import { endPoint } from "@/constants";
import useSWR from "swr";
import { getToken } from "../auth";

export default function getAllImportNote({
  page,
  limit,
  filterString,
  token,
}: {
  page: string;
  limit?: number;
  filterString?: string;
  token: string;
}) {
  const fetcher = (url: string) => {
    return fetch(url, {
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
          paging: json.paging,
          data: json.data,
        };
      });
  };
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
