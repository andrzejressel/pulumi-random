// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.random.RandomStringArgs;
import com.pulumi.random.Utilities;
import com.pulumi.random.inputs.RandomStringState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The resource `random.RandomString` generates a random permutation of alphanumeric characters and optionally special characters.
 * 
 * This resource *does* use a cryptographic random number generator.
 * 
 * Historically this resource&#39;s intended usage has been ambiguous as the original example used it in a password. For backwards compatibility it will continue to exist. For unique ids please use random_id, for sensitive random values please use random_password.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomString;
 * import com.pulumi.random.RandomStringArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var random = new RandomString(&#34;random&#34;, RandomStringArgs.builder()        
 *             .length(16)
 *             .overrideSpecial(&#34;/@£$&#34;)
 *             .special(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * You can import external strings into your Pulumi programs as RandomString resources as follows:
 * 
 * ```sh
 *  $ import random:index/randomString:RandomString newString myspecialdata
 * ```
 * 
 * This command will encode the `myspecialdata` token in Pulumi state and generate a code suggestion to include a new RandomString resource in your Pulumi program. Include the suggested code and do a `pulumi up`. Your data is now stored in Pulumi, and you can reference it in your Pulumi program as `newString.result`.
 * 
 * If the data needs to be stored securily as a secret, consider using the RandomPassword resource instead.
 * 
 */
@ResourceType(type="random:index/randomString:RandomString")
public class RandomString extends com.pulumi.resources.CustomResource {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    @Export(name="keepers", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    public Output<Optional<Map<String,String>>> keepers() {
        return Codegen.optional(this.keepers);
    }
    /**
     * The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
     * 
     */
    @Export(name="length", refs={Integer.class}, tree="[0]")
    private Output<Integer> length;

    /**
     * @return The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` + `min_lower` + `min_numeric` + `min_special`).
     * 
     */
    public Output<Integer> length() {
        return this.length;
    }
    /**
     * Include lowercase alphabet characters in the result. Default value is `true`.
     * 
     */
    @Export(name="lower", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> lower;

    /**
     * @return Include lowercase alphabet characters in the result. Default value is `true`.
     * 
     */
    public Output<Boolean> lower() {
        return this.lower;
    }
    /**
     * Minimum number of lowercase alphabet characters in the result. Default value is `0`.
     * 
     */
    @Export(name="minLower", refs={Integer.class}, tree="[0]")
    private Output<Integer> minLower;

    /**
     * @return Minimum number of lowercase alphabet characters in the result. Default value is `0`.
     * 
     */
    public Output<Integer> minLower() {
        return this.minLower;
    }
    /**
     * Minimum number of numeric characters in the result. Default value is `0`.
     * 
     */
    @Export(name="minNumeric", refs={Integer.class}, tree="[0]")
    private Output<Integer> minNumeric;

    /**
     * @return Minimum number of numeric characters in the result. Default value is `0`.
     * 
     */
    public Output<Integer> minNumeric() {
        return this.minNumeric;
    }
    /**
     * Minimum number of special characters in the result. Default value is `0`.
     * 
     */
    @Export(name="minSpecial", refs={Integer.class}, tree="[0]")
    private Output<Integer> minSpecial;

    /**
     * @return Minimum number of special characters in the result. Default value is `0`.
     * 
     */
    public Output<Integer> minSpecial() {
        return this.minSpecial;
    }
    /**
     * Minimum number of uppercase alphabet characters in the result. Default value is `0`.
     * 
     */
    @Export(name="minUpper", refs={Integer.class}, tree="[0]")
    private Output<Integer> minUpper;

    /**
     * @return Minimum number of uppercase alphabet characters in the result. Default value is `0`.
     * 
     */
    public Output<Integer> minUpper() {
        return this.minUpper;
    }
    /**
     * Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
     * 
     * @deprecated
     * **NOTE**: This is deprecated, use `numeric` instead.
     * 
     */
    @Deprecated /* **NOTE**: This is deprecated, use `numeric` instead. */
    @Export(name="number", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> number;

    /**
     * @return Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
     * 
     */
    public Output<Boolean> number() {
        return this.number;
    }
    /**
     * Include numeric characters in the result. Default value is `true`.
     * 
     */
    @Export(name="numeric", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> numeric;

    /**
     * @return Include numeric characters in the result. Default value is `true`.
     * 
     */
    public Output<Boolean> numeric() {
        return this.numeric;
    }
    /**
     * Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     * 
     */
    @Export(name="overrideSpecial", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> overrideSpecial;

    /**
     * @return Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     * 
     */
    public Output<Optional<String>> overrideSpecial() {
        return Codegen.optional(this.overrideSpecial);
    }
    /**
     * The generated random string.
     * 
     */
    @Export(name="result", refs={String.class}, tree="[0]")
    private Output<String> result;

    /**
     * @return The generated random string.
     * 
     */
    public Output<String> result() {
        return this.result;
    }
    /**
     * Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
     * 
     */
    @Export(name="special", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> special;

    /**
     * @return Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
     * 
     */
    public Output<Boolean> special() {
        return this.special;
    }
    /**
     * Include uppercase alphabet characters in the result. Default value is `true`.
     * 
     */
    @Export(name="upper", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> upper;

    /**
     * @return Include uppercase alphabet characters in the result. Default value is `true`.
     * 
     */
    public Output<Boolean> upper() {
        return this.upper;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RandomString(String name) {
        this(name, RandomStringArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RandomString(String name, RandomStringArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RandomString(String name, RandomStringArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomString:RandomString", name, args == null ? RandomStringArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RandomString(String name, Output<String> id, @Nullable RandomStringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("random:index/randomString:RandomString", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RandomString get(String name, Output<String> id, @Nullable RandomStringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RandomString(name, id, state, options);
    }
}
