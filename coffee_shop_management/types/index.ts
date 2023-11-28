import { IconType } from "react-icons";

export type Product = {
  id: string;
  name: string;
  price: number;
  status: "active" | "inactive";
  image?: string;
  idCate?: string;
};

export type SidebarItem = {
  title: string;
  href: string;
  icon?: IconType;
  submenu?: boolean;
  subMenuItems?: SidebarItem[];
};

export type Category = {
  id: string;
  name: string;
  quantity: number;
};

export type MeasureUnit = {
  id: string;
  name: "g" | "kg" | "l" | "ml" | "đơn vị";
  covertDetails?: {
    measureUnit: MeasureUnit;
    times: number;
  };
};

export type IngredientForChoose = {
  id: string;
  name: string;
  unitId: string;
};

export type Ingredient = {
  id: string;
  name: string;
  total: number;
  unit: MeasureUnit;
  price: number;
};
export type IngredientDetail = {
  idIngre: string;
  expirationDate: Date;
  quantity: number;
};

export type Staff = {
  id: string;
  name: string;
  role: string;
};
export type Role = {
  id: string;
  name: string;
  function?: string[];
};
export type RoleFunction = {
  id: string;
  name: string;
};

export enum StatusString {
  Inprogress = "Đang xử lý",
  Done = "Đã nhập",
  Cancel = "Đã huỷ",
}
export type ImportNote = {
  id: string;
  supplierId: string;
  totalPrice: number;
  status: StatusString;
  createBy: string;
  closeBy?: string;
  createAt: Date;
  closeAt?: Date;
};
export type ExportNote = {
  id: string;
  createBy: string;
  createAt: Date;
  reason: string;
};

export type Customer = {
  id: string;
  name: string;
  email?: string;
  phone: string;
  point: number;
};

export interface UnitListProps {
  unit: string;
  setUnit: (unit: string) => void;
}
export interface RoleListProps {
  role: string;
  setRole: (role: string) => void;
}

export interface CategoryListProps {
  category: string;
  setCategory: (category: string) => void;
  canAdd?: boolean;
}
