import { apiKey, endPoint } from "@/constants";
import { Staff } from "@/types";
import useSWR from "swr";

export default function getAllStaff(token: string) {
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
      .then((json) => json.data);
  const { data, error, isLoading } = useSWR(`${endPoint}/users/all`, fetcher);

  return {
    staffs: data as Staff[],
    isLoading,
    isError: error,
  };
}
