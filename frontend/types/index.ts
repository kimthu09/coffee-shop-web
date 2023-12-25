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
};

export type IngredientForChoose = {
  id: string;
  name: string;
  unitId: string;
};

export type Ingredient = {
  id: string;
  name: string;
  amount: number;
  price: number;
  measureType: string;
};
export type IngredientDetail = {
  idIngre: string;
  expirationDate: Date;
  quantity: number;
};

export type Staff = {
  id: string;
  name: string;
  email: string;
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

export type ImportNote = {
  id: string;
  supplierId: string;
  totalPrice: number;
  status: StatusNote;
  closedAt?: Date;
  closedBy?: {
    id: string;
    name: string;
  };
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
  supplier: {
    id: string;
    name: string;
    phone: string;
  };
};
export type ImportNoteDetail = {
  ingredient: {
    id: string;
    name: string;
    measureType: string;
  };
  price: number;
  amountImport: number;
};
export type ExportNote = {
  id: string;
  createBy: string;
  createAt: Date;
  reason: string;
};
export enum StatusNote {
  Inprogress = "InProgress",
  Done = "Done",
  Cancel = "Cancel",
}
export type Customer = {
  id: string;
  name: string;
  email?: string;
  phone: string;
  point: number;
};

export type Supplier = {
  id: string;
  name: string;
  email?: string;
  phone: string;
  debt: number;
};
export interface UnitListProps {
  unit: string;
  setUnit: (unit: string) => void;
}
export interface RoleListProps {
  role: string;
  setRole: (role: string) => void;
}
export type FilterValue = {
  filters: {
    type: string;
    value: string;
  }[];
};
export interface StaffListProps {
  staff: string;
  setStaff: (role: string) => void;
}
export interface CategoryListProps {
  checkedCategory: Array<string>;
  onCheckChanged: (idCate: string) => void;
  canAdd?: boolean;
}
export type PagingProps = {
  page: number;
  limit: number;
  total: number;
};
