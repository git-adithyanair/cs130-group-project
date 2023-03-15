import React from "react";
import { render, fireEvent } from "@testing-library/react-native";
import ActiveItemCard from "../app/components/ActiveItemCard";

describe("active errand test suite", () => {
  test("given item found has no status, clicking found displays Item Found", () => {
    const mockOnPressUpdateFound = jest.fn();

    const component = render(
      <ActiveItemCard
        item={{
          found: { Valid: false, Bool: false },
          preferred_brand: { Valid: false },
          extra_notes: { Valid: false },
          image: { Valid: false },
        }}
        onPressUpdateFound={mockOnPressUpdateFound}
      />
    );

    fireEvent.press(component.queryByText("Found"));

    expect(mockOnPressUpdateFound).toHaveBeenCalled();
    expect(component.findByText("Item Found")).not.toBeNull();
  });

  test("given item found has no status, clicking not found displays Item Not Found", () => {
    const mockOnPressUpdateFound = jest.fn();

    const component = render(
      <ActiveItemCard
        item={{
          found: { Valid: false, Bool: false },
          preferred_brand: { Valid: false },
          extra_notes: { Valid: false },
          image: { Valid: false },
        }}
        onPressUpdateFound={mockOnPressUpdateFound}
      />
    );

    fireEvent.press(component.queryByText("Not Found"));

    expect(mockOnPressUpdateFound).toHaveBeenCalled();
    expect(component.findByText("Item Not Found")).not.toBeNull();
  });

  test("given item has already been found, buttons are hidden", () => {
    const mockOnPressUpdateFound = jest.fn();

    const component = render(
      <ActiveItemCard
        item={{
          found: { Valid: true, Bool: true },
          preferred_brand: { Valid: false },
          extra_notes: { Valid: false },
          image: { Valid: false },
        }}
        onPressUpdateFound={mockOnPressUpdateFound}
      />
    );

    expect(component.queryByText("Item Found")).not.toBeNull();
    expect(component.queryByText("Found")).toBeNull();
  });

  test("given item with valid arguments, expect them to be displayed on screen", () => {
    const mockOnPressUpdateFound = jest.fn();
    const itemData = {
      found: { Valid: true, Bool: true },
      name: "This is the name",
      quantity: 3,
      quantity_type: "Type",
      preferred_brand: { Valid: false },
      extra_notes: { Valid: true, String: "Test extra notes" },
      image: { Valid: false },
    };

    const component = render(
      <ActiveItemCard
        item={itemData}
        onPressUpdateFound={mockOnPressUpdateFound}
      />
    );

    expect(component.queryByText(itemData.name)).not.toBeNull();
    expect(
      component.queryByText(
        `Amount: ${itemData.quantity} (${itemData.quantity_type})`
      )
    ).not.toBeNull();
    expect(component.queryByText("Item Found")).not.toBeNull();
    expect(component.queryByText("Preferred Brand: Any")).not.toBeNull();
    expect(
      component.queryByText(`Extra Notes: ${itemData.extra_notes.String}`)
    ).not.toBeNull();

    expect(component.queryByText("Found")).toBeNull();
    expect(component.queryByText("Not Found")).toBeNull();
  });
});
