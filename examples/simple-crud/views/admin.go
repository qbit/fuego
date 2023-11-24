package views

import (
	"simple-crud/store"

	"github.com/go-fuego/fuego"
)

func (rs Ressource) pageAdmin(c fuego.Ctx[any]) (any, error) {
	return c.Redirect(301, "/admin/recipes")
}

func (rs Ressource) deleteRecipe(c fuego.Ctx[any]) (any, error) {
	id := c.QueryParam("id") // TODO use PathParam
	err := rs.Queries.DeleteRecipe(c.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.Redirect(301, "/admin/recipes")
}

func (rs Ressource) adminRecipes(c fuego.Ctx[any]) (fuego.HTML, error) {
	recipes, err := rs.Queries.GetRecipes(c.Context())
	if err != nil {
		return "", err
	}

	return c.Render("pages/admin/recipes.page.html", fuego.H{
		"Recipes": recipes,
	})
}

func (rs Ressource) adminOneRecipe(c fuego.Ctx[any]) (fuego.HTML, error) {
	id := c.QueryParam("id") // TODO use PathParam

	recipe, err := rs.Queries.GetRecipe(c.Context(), id)
	if err != nil {
		return "", err
	}

	ingredients, err := rs.Queries.GetIngredientsOfRecipe(c.Context(), id)
	if err != nil {
		return "", err
	}

	return c.Render("pages/admin/single-recipe.page.html", fuego.H{
		"Name":         recipe.Name,
		"Description":  recipe.Description,
		"Ingredients":  ingredients,
		"Instructions": nil,
	})
}

func (rs Ressource) adminAddRecipes(c fuego.Ctx[store.CreateRecipeParams]) (any, error) {
	body, err := c.Body()
	if err != nil {
		return "", err
	}

	_, err = rs.Queries.CreateRecipe(c.Context(), body)
	if err != nil {
		return "", err
	}

	return c.Redirect(301, "/admin/recipes")
}

func (rs Ressource) adminIngredients(c fuego.Ctx[any]) (fuego.HTML, error) {
	ingredients, err := rs.Queries.GetIngredients(c.Context())
	if err != nil {
		return "", err
	}

	return c.Render("pages/admin/ingredients.page.html", fuego.H{
		"Ingredients": ingredients,
	})
}

func (rs Ressource) adminAddIngredient(c fuego.Ctx[store.CreateIngredientParams]) (any, error) {
	body, err := c.Body()
	if err != nil {
		return "", err
	}

	_, err = rs.Queries.CreateIngredient(c.Context(), body)
	if err != nil {
		return "", err
	}

	return c.Redirect(301, "/admin/ingredients")
}