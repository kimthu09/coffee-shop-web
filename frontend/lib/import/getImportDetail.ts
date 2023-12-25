import { apiKey, endPoint } from "@/constants";
import useSWR from "swr";

const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
    cache: "no-store",
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => json.data);

export default function getImportNoteDetail(idNote: string) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/importNotes/${idNote}`,
    fetcher
  );

  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
