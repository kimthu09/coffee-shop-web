import { MeasureType, StatusNote } from "@/types";
import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const toVND = (money: number) => {
  const formatted = new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(money);
  return formatted;
};

export const removeAccents = (str: string) => {
  return str
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/đ/g, "d")
    .replace(/Đ/g, "D");
};

export const toUnit = (str: string) => {
  if (str === MeasureType.Volume) {
    return "ml";
  } else if (str === MeasureType.Weight) {
    return "g";
  } else {
    return "đơn vị";
  }
};

export const statusNoteToString = (status: StatusNote) => {
  if (status === StatusNote.Inprogress) {
    return "Đang tiến hành";
  } else if (status === StatusNote.Done) {
    return "Đã hoàn thành";
  } else {
    return "Đã hủy";
  }
};
